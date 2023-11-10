// Initialize web3
if (window.ethereum) {
    window.web3 = new Web3(ethereum);
    try {
        // Request account access if needed
        ethereum.enable();
    } catch (error) {
        // User denied account access...
        console.error("User denied account access");
    }
} else if (window.web3) {
    window.web3 = new Web3(web3.currentProvider);
} else {
    console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
}

// Set up event listener for the send order form
document.getElementById('sendOrderForm').addEventListener('submit', function(event) {
    event.preventDefault();
    // Call the createOrder function of your smart contract here
});

// Set up event listener for the track order form
document.getElementById('trackOrderForm').addEventListener('submit', function(event) {
    event.preventDefault();
    // Call the getOrder function of your smart contract here
});

