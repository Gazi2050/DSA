// Step 1: Define the fetchData function that returns a Promise
function fetchData() {
    return new Promise((resolve, reject) => {
        console.log("Fetching data...");

        // Step 2: Use setTimeout to simulate an async operation with a 1.5-second delay
        setTimeout(() => {
            const success = true; // Change this to 'false' to test error handling

            if (success) {
                const data = { id: 1, name: "User" }; // Simulated fetched data
                console.log("Data fetched successfully");

                // Step 3: Resolve the promise with the fetched data
                resolve(data);
            } else {
                // Step 4: Reject the promise with an error message if success is false
                reject("Failed to fetch data");
            }
        }, 1500);
    });
}

// Step 5: Call fetchData, handling the resolved and rejected states with .then() and .catch()
fetchData()
    .then((data) => {
        // Step 6: Process the data if the promise is resolved
        console.log(`Processing data for ${data.name}`);
    })
    .catch((error) => {
        // Step 7: Handle any errors if the promise is rejected
        console.log(error);
    });
