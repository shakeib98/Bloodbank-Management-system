// sum.rs

fn sum(numbers: &[i32]) -> i32 {
    numbers.iter().sum()
}

fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    let total = sum(&numbers);
    println!("The sum of {:?} is {}", numbers, total);
}