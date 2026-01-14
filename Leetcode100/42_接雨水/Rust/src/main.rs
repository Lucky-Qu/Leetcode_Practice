//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

struct Solution;

impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut left: usize = 0;
        let mut right = height.len() - 1;
        let mut left_max = height[left];
        let mut right_max = height[right];
        let mut sum = 0;
        while left < right {
            // 移动短板
            if left_max < right_max {
                left += 1;
                if height[left] > left_max {
                    left_max = height[left];
                }else {
                    sum += left_max - height[left];
                }
            }else {
                right -= 1;
                if height[right] > right_max {
                    right_max = height[right];
                }else {
                    sum += right_max - height[right];
                }
            }
        }
        sum
    }
}

fn main() {
    println!("{}", Solution::trap(vec![0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
}