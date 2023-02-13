use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use serde_json::json;
use validator::Validate;
use futures::{StreamExt, TryStreamExt};
use k8s_openapi::apiextensions_apiserver::pkg::apis::apiextensions::v1::CustomResourceDefinition;
use kube::{
    api::{Api, DeleteParams, ListParams, PatchParams, Patch, ResourceExt},
    core::CustomResourceExt,
    Client, CustomResource,
    runtime::{watcher, WatchStreamExt, wait::{conditions, await_condition}},
};

// Our custom resource
#[derive(CustomResource, Deserialize, Serialize, Clone, Debug, Validate, JsonSchema)]
#[kube(group = "clux.dev", version = "v1", kind = "Foo", namespaced)]
pub struct FooSpec {
    info: String,
    #[validate(length(min = 3))]
    name: String,
    replicas: i32,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::try_default().await?;
    let crds: Api<CustomResourceDefinition> = Api::all(client.clone());

    // Apply the CRD so users can create Foo instances in Kubernetes
    crds.patch("foos.clux.dev",
               &PatchParams::apply("my_manager"),
               &Patch::Apply(Foo::crd())
    ).await?;

    // Wait for the CRD to be ready
    tokio::time::timeout(
        std::time::Duration::from_secs(10),
        await_condition(crds, "foos.clux.dev", conditions::is_crd_established())
    ).await?;

    // Watch for changes to foos in the configured namespace
    let foos: Api<Foo> = Api::default_namespaced(client.clone());
    let lp = ListParams::default();
    let mut apply_stream = watcher(foos, lp).applied_objects().boxed();
    while let Some(f) = apply_stream.try_next().await? {
        println!("saw apply to {}", f.name_any());
    }
    Ok(())
}