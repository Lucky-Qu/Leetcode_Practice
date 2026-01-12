
struct Solution {}

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut left: i32 = 0;
        let mut right: i32 = (height.len() - 1) as i32;
        let mut area:i32 = 0;
        let mut max_area:i32 = 0;
        loop {
            // 计算并比较面积
            area = (right - left) * std::cmp::min(height[left as usize], height[right as usize]);
            if area > max_area {
                max_area = area
            };
            // 移动指针
            if height[left as usize] < height[right as usize] {
                left += 1;
            }else {
                right -= 1;
            };
            // 边界检查
            if left >= right {
                break max_area;
            }
        }
    }
}

fn main() {
    println!("{}", Solution::max_area(vec![1,8,6,2,5,4,8,3,7]));
}