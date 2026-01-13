use std::env::var;

// 给你一个整数数组 nums
// 判断是否存在三元组 [nums[i], nums[j], nums[k]]
// 满足 i != j、i != k 且 j != k
// 同时还满足 nums[i] + nums[j] + nums[k] == 0
// 请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
struct Solution {}

impl Solution {
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        let n = nums.len();
        let mut res = Vec::new();

        if n < 3 {
            return res;
        }

        nums.sort();

        let mut i = 0;
        while i < n {
            let first = nums[i];

            // 剪枝
            if first > 0 {
                break;
            }

            // 外层去重
            if i > 0 && nums[i] == nums[i - 1] {
                i += 1;
                continue;
            }

            let mut left = i + 1;
            let mut right = n - 1;

            while left < right {
                let sum = first + nums[left] + nums[right];

                if sum == 0 {
                    res.push(vec![first, nums[left], nums[right]]);

                    // 内层去重：保存旧值
                    let lv = nums[left];
                    let rv = nums[right];

                    while left < right && nums[left] == lv {
                        left += 1;
                    }
                    while left < right && nums[right] == rv {
                        right -= 1;
                    }
                } else if sum < 0 {
                    left += 1;
                } else {
                    right -= 1;
                }
            }

            i += 1;
        }

        res
    }
}
fn main() {
    println!("{:?}", Solution::three_sum(vec![-1,0,1,2,-1,-4]))
}