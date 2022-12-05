## Informer



Informer 负责与 kubernetes APIServer 进行 Watch 操作，Watch 的资源，可以是 kubernetes 内置资源对象，也可以 CRD。

Informer 是一个带有本地缓存以及索引机制的核心工具包，当请求为 查询 操作的时候，会优先从本地缓存内存去查询数据，而 创建、更新、删除这类操作，则会根据事件通知写入队列 DeltaFIFO 中，同时对应的事件处理过后，更新本地缓存，使得本地缓存与 ETCD 的数据保持一致性。

Informer 抽象出来的这个缓存层，将 查询 相关操作压力接收了下来，这样就不必每次都去调用 APIService 的接口，减轻了 APIServer 的数据交互压力。

Informer 相关组件的概念：

* Reflector：使用 List-Watch 来保证本地缓存数据的准确性和顺序一致性的，List 对应资源的全量列表数据，Watch 负责变化部分的数据，Watch 指定的 Kubernetes 资源，当 Watch 的资源发送变化时，触发变更的事件，如：Added、Updated及Deleted 事件，并将资源对象的变化事件存到本地队列 DeltaFIFO 中。
* DeltaFIFO：是一个增量的队列，记录了资源变化过程，Reflector 就相当于队列的生产者。这个组件可以拆分为两个部分来理解。FIFO 就是一个队列，拥有队列基本操作方法如ADD、UPDATE、DELETE、LIST、POP及CLOSE等，而Delta是一个资源对象存储，保存存储对象的消息类型如Added、Updated及Deleted 等。
* Indexer：是用来存储资源对象并自带索引功能的本地存储，Reflector 从 DeltaFIFO 中消费出来的资源对象存储到 Indexer，Indexer 与 Etcd 中的数据是完全一致的，从而client-go 可以本地读取，这样来减少 k8s APIServer 的资源交互压力。



### List-Watch

List-Watch 机制是 k8s 中的异步消息通知机制，通过它能有效的确保了消息的实时性、顺序性及可靠性。

List-Watch 分为两个部门：

* List 负责调用资源对应的 k8s APIServer 的 RESTful API获取全局数据列表，并且同步到（本地缓存中）DeltaFIFO。
* Watch 负责监听资源的变化，并调用相应事件处理函数进行处理，同时维护更新本地缓存，使得本地队列与 Etcd 总数据保持一致。

List 是基于 HTTP 中的短链接实现，Watch 则是基于 HTTP 长链接实现的，Watch 使用长链接方式是为了减轻 k8s APIServer 的访问压力。

### Reflector

Reflector 是client-go 中用来监听指定资源的组件，当资源发生变化的时候如ADDED、UPDATE及DELETE操作时，会以事件的形式存入**本地队列**，然后有对应的方法处理。

在 Reflector 中，核心的部分分别是 List-Watch，其他功能基本都是围绕它来搞的。

在实例化 Reflector 的过程中，其中有ListerWatcher的接口对象，这个结构体对象有两个方法，分别是List和Watch这两个方法实现List-Watch功能。

Reflector 核心逻辑：

* List：调用List方法获取资源全部列表数据，转换为资源对象列表，然后保存到本地缓存中，
* 定时同步：定时器定时触发同步机制，定时更新缓存数据，在Reflector结构体对象中，是可以配置定时同步的周期时间的。
* Watch：监听资源的变化，并且调用对应的事件处理函数来处理

Reflector 组件对于数据更新同步，都是基于 ResourceVersion 来进行，每个资源对象都会有 ResourceVersion 这个属性，当数据发生变化时候 ResourceVersion 也会将以递增的形式更新，这样就确保事件的更新顺序了。



### DeltaFIFO

DeltaFIFO 是一个增量的本地队列，记录资源对象的变化过程。它生产者是 Reflector 组件，将监听到的对象，同步到 DeltaFIFO中。

* FIFO 是一个先进先出的本地队列，负责接收 Reflector 传递过来的事件，并将其按照顺序存储，然后等待事件的处理，若同时出现多个相同的事件，则会被处理一次。FIFO 是队列那么它也会拥有队列相关操作的方法，可以通过 Queue 这个接口对象来实现队列需求的方法的，同时还根据需要扩展一些其他的方法如Pop，AddIfNotPresent等。
* Delta 是一个资源对象的存储，有两个属性分别是 Type 和 Object
    * Type 表示这个事件类型，如Added表示增量，Update 表示更新等。
    * Object 是一个接口类型，它表示k8s 资源对象如Pod和Service等。





