
struct Solution;

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut m = std::collections::HashMap::<u8, usize>::new();
        let mut left = 0;
        let mut right = 0;
        let mut max_len = 0;
        let s = s.as_bytes();
        while right < s.len() {
            match m.get(&s[right]){
                Some(&index) => {
                    if index >= left {
                        left = index + 1;
                    }
                }
                None => {
                }
            }
            m.insert(s[right], right);
            max_len = std::cmp::max(max_len, (right + 1 - left) as i32);
            right += 1;
        }
        max_len
    }
}

fn main() {
    println!("{}", Solution::length_of_longest_substring("pwwkew".to_string()))
}
