// Step 1: Define the fetchData function that accepts a callback parameter
function fetchData(callback) {
    console.log("Fetching data...");

    // Step 2: Use setTimeout to simulate data fetching that takes 1.5 seconds
    setTimeout(() => {
        const data = { id: 1, name: "User" }; // Simulated fetched data
        console.log("Data fetched successfully");

        // Step 3: Call the callback function and pass the fetched data to it
        callback(data);
    }, 1500);
}

// Step 4: Define the processData function that takes data as an argument
function processData(data) {
    console.log(`Processing data for ${data.name}`); // Process the fetched data
}

// Step 5: Call fetchData and pass processData as the callback function
fetchData(processData);
