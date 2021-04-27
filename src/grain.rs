pub fn grain() -> i128 {
    let mut s: i128 = 0;
    for i in 0..64 {
        s += 2i128.pow(i);
    }
    return s;
}

#[test]
fn test_grain() {
    assert_eq!(18446744073709551615, grain());
}
