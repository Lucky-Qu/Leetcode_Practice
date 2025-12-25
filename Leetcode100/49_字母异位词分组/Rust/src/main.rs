// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 解释：
// 在 strs 中没有字符串可以通过重新排列来形成 "bat"。
// 字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
// 字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
// 提示：
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
struct Solution {}

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut map: std::collections::HashMap<[i32; 26], Vec<String>> = std::collections::HashMap::new();

        // 1. 消费 strs
        for s in strs {
            let mut key = [0i32; 26];

            // 2. byte 遍历
            for b in s.as_bytes() {
                key[(b - b'a') as usize] += 1;
            }

            // 3. entry + push
            map.entry(key)
                .or_insert_with(Vec::new)
                .push(s);
        }

        // 4. 直接 move 出 map 的 value（无 clone）
        map.into_values().collect()
    }
}

fn main() {
    let result = Solution::group_anagrams(std::vec!["bat".to_string(), "tab".to_string(), "save".to_string()]);
    println!("{:?}",result)

}