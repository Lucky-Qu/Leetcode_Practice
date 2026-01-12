// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 示例 2:
// 输入: nums = [0]
// 输出: [0]

struct Solution {}

impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut left = 0;
        let mut right = 0;
        loop {
            if right >= nums.len() {
                break;
            }

            if nums[left] != 0 {
                left += 1;
                right +=1;
                continue;
            }
            if nums[right] != 0 {
                nums.swap(left, right);
                left += 1;
            }
            right += 1;
        }
    }
}

fn main() {
    let mut result = vec![0,1,0,3,12];
    Solution::move_zeroes(&mut result);
    println!("{:?}", result);
}
