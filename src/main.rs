mod clock;

fn main() {
    let mut t = clock::Timer::new();
    t.add(1000);
    println!("{}", t);
}
