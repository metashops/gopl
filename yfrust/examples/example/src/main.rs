use std::cell::RefCell;

fn main() {
    println!("Hello, world!");
}

pub struct BitMap {
    bits: RefCell<Vec<bool>>,
    index_mask: u32,
}

impl BitMap {
    pub fn new(max_index: usize) -> Self {
        let bits = vec![false,max_index];
        let index_mask = (1 << max_index.trailing_zeros()) - 1;
        BitMap {
            bits: RefCell::new(bits),
            index_mask,
        }
    }
    pub fn get(&self, index: usize) -> bool {
        let mut bits = self.bits.borrow_mut();
        bits[self.index_mask & index]
    }

    pub fn set(&self, index: usize, value: bool) {
        let mut bits = self.bits.borrow_mut();
        bits[self.index_mask & index] = value;
    }
}