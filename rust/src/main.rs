use algonaut_transaction::account::Account;
use std::time::Instant;

fn main() {
    let mut count = 0;
    let mut addresses: [String; 5] = Default::default();
    let t0 = Instant::now();
    while count < 5 {
        let account = Account::generate().address().to_string();
        if account.starts_with("TEST") {
            addresses[count] = account;
            count += 1
        }
    }
    println!(
        "Rust found 5 addresses in {:?} secs",
        t0.elapsed().as_secs()
    );
    println!("{:?}", addresses);
}
