// Step 1: Define a function that returns a promise (simulating data fetching)
function fetchData(data) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(`Fetched data: ${data}`);
        }, 2000); // Simulate a 2-second network request delay
    });
}

// Step 2: Define an async function to use async/await syntax
async function getData() {
    console.log("Starting data fetch...");

    // Step 3: Use 'await' to wait for the promise to resolve
    // 'await' will pause execution here until fetchData resolves
    const result = await fetchData("Sample Data");

    // Step 4: Once fetchData resolves, execution resumes here
    console.log(result);

    // Step 5: Continue executing other code after awaiting
    console.log("Data fetch completed.");
}

// Step 6: Call the async function to see everything in action
getData();
