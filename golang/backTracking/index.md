解决一个回溯问题，实际上就是一个决策树的遍历过程，站在回溯树的一个节点上，你只需要思考 3 个问题：

1、路径：也就是已经做出的选择。

2、选择列表：也就是你当前可以做的选择。

3、结束条件：也就是到达决策树底层，无法再做选择的条件。

代码方面，回溯算法的框架：

```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

https://labuladong.github.io/algo/1/9/

## **总结**
- 来回顾一下排列/组合/子集问题的三种形式在代码上的区别。

- 由于子集问题和组合问题本质上是一样的，无非就是 base case 有一些区别，所以把这两个问题放在一起看。

- 形式一、元素无重不可复选，即 nums 中的元素都是唯一的，每个元素最多只能被使用一次，backtrack 核心代码如下：

    ```java
    /* 组合/子集问题回溯算法框架 */
    void backtrack(int[] nums, int start) {
        // 回溯算法标准框架
        for (int i = start; i < nums.length; i++) {
            // 做选择
            track.addLast(nums[i]);
            // 注意参数
            backtrack(nums, i + 1);
            // 撤销选择
            track.removeLast();
        }
    }
    
    /* 排列问题回溯算法框架 */
    void backtrack(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            // 剪枝逻辑
            if (used[i]) {
                continue;
            }
            // 做选择
            used[i] = true;
            track.addLast(nums[i]);
    
            backtrack(nums);
            // 撤销选择
            track.removeLast();
            used[i] = false;
        }
    }
    ```

- 形式二、元素可重不可复选，即 nums 中的元素可以存在重复，每个元素最多只能被使用一次，其关键在于`排序和剪枝`，backtrack 核心代码如下：

    ```java
    Arrays.sort(nums);
    /* 组合/子集问题回溯算法框架 */
    void backtrack(int[] nums, int start) {
        // 回溯算法标准框架
        for (int i = start; i < nums.length; i++) {
            // 剪枝逻辑，跳过值相同的相邻树枝
            if (i > start && nums[i] == nums[i - 1]) {
                continue;
            }
            // 做选择
            track.addLast(nums[i]);
            // 注意参数
            backtrack(nums, i + 1);
            // 撤销选择
            track.removeLast();
        }
    }
    
    
    Arrays.sort(nums);
    /* 排列问题回溯算法框架 */
    void backtrack(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            // 剪枝逻辑
            if (used[i]) {
                continue;
            }
            // 剪枝逻辑，固定相同的元素在排列中的相对位置
            if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) {
                continue;
            }
            // 做选择
            used[i] = true;
            track.addLast(nums[i]);
    
            backtrack(nums);
            // 撤销选择
            track.removeLast();
            used[i] = false;
        }
    }
    ```

- 形式三、元素无重可复选，即 nums 中的元素都是唯一的，每个元素可以被使用若干次，只要删掉去重逻辑即可，backtrack 核心代码如下：

    ```java
    /* 组合/子集问题回溯算法框架 */
    void backtrack(int[] nums, int start) {
        // 回溯算法标准框架
        for (int i = start; i < nums.length; i++) {
            // 做选择
            track.addLast(nums[i]);
            // 注意参数
            backtrack(nums, i);
            // 撤销选择
            track.removeLast();
        }
    }
    
    
    /* 排列问题回溯算法框架 */
    void backtrack(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            // 做选择
            track.addLast(nums[i]);
            backtrack(nums);
            // 撤销选择
            track.removeLast();
        }
    }
    ```

- **只要从树的角度思考，这些问题看似复杂多变，实则改改 base case 就能解决**