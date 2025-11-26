
struct Solution;
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
//你可以按任意顺序返回答案。
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = std::collections::HashMap::<i32, i32>::new();
        for (i, v) in nums.iter().enumerate() {
            if map.contains_key(&(target - v)) {
                return Vec::from([i as i32, *map.get(&(target - v)).unwrap()])
            }
            map.insert(*v, i as i32 );
        }
        Vec::from([-1, -1])
    }
}

fn main() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    let result = Solution::two_sum(nums, target);
    println!("{:?}", result);
}
