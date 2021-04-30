use std::fmt::Debug;
use std::fmt::Display;
use std::fmt::Formatter;
use std::fmt::Result;

pub struct Timer {
    minute: i64,
}

impl Timer {
    pub fn new() -> Timer {
        return Timer { minute: 0 };
    }

    pub fn hour(&self) -> i64 {
        return self.minute / 60;
    }

    pub fn minute(&self) -> i64 {
        return self.minute % 60;
    }

    pub fn add(&mut self, minute: i64){
        self.minute += minute;
        self.minute %= 60 * 24;
    }
}

impl Display for Timer {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        let h = self.hour();
        let m = self.minute();
        return write!(f, "{:>02}:{:>02}", h, m);
    }
}

impl PartialEq for Timer {
    fn eq(&self, another: &Timer) -> bool {
        return self.hour() == another.hour() && self.minute() == another.minute();
    }
}

impl Debug for Timer {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        return write!(f, "{:>02}:{:>02}", self.hour(), self.minute());
    }
}

#[test]
fn test_equal(){
    let mut t0 = Timer::new();
    let mut t1 = Timer::new();
    t0.add(1000);
    t1.add(300);
    assert_ne!(t0, t1);
    t1.add(700);
    assert_eq!(t0, t1);
}