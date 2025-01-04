Using `testify/mock`, you can make detailed assertions about how many times a mocked method is called, among other things. This is especially useful to ensure that your functions are not only tested for returned values but also called the expected number of times during execution.

Here’s how to enhance the previous sample with additional assertions, including checking the number of times the mocked functions are called.

### Example with Mock Assertions

We will add assertions to check that certain methods are called exactly once and specify any parameters if needed.

#### Mock Repository

First, we will set up our mock repository as before, but this time we'll add assertions for call counts.

### Explanation of Mock Assertions

1. **Expectations Setup**:
   - The `.Once()` method specifies that the mocked method `GetUserByID` is expected to be called exactly once during the test.

2. **Assertions**:
   - `mockRepo.AssertCalled(t, "GetUserByID", 1)`: Asserts that `GetUserByID` was called with the parameter `1`.
   - In the second test (`TestGetUser_NotFound`), it checks that the `GetUserByID` method is called with `999`.

3. **Assert Not Called**:
   - In the `TestGetUser_InvalidID`, we assert that the repository's `GetUserByID` method is not called when the user ID is invalid (e.g., `0`). This demonstrates how to use `AssertNotCalled` to ensure that certain methods