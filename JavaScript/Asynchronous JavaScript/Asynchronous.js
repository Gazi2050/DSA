// Step 1: Start asynchronous task 1
console.log("Starting asynchronous task 1");

// Step 2: Use setTimeout to delay the next console log by 2 seconds (2000 ms)
setTimeout(() => {
    console.log("Task 1 completed after 2 seconds"); // This line executes after 2 seconds
}, 2000);

// Step 3: Print task 2 completion immediately without waiting for the setTimeout
console.log("Task 2 completed immediately");
