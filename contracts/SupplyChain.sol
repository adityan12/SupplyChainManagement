// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SupplyChain {
    // Define an Order struct with relevant fields
    struct Order {
        uint256 orderNumber;
        string name;
        string email;
        string phoneNumber;
        string fromAddress;
        string toAddress;
        bool isShipped;
    }

    // State variable to keep track of the last order number
    uint256 private lastOrderNumber = 0;

    // Mapping from order number to Order struct
    mapping(uint256 => Order) public orders;

    // Event to emit when a new order is created
    event OrderCreated(uint256 orderNumber);

    // Function to create a new order
    function createOrder(
        string memory _name,
        string memory _email,
        string memory _phoneNumber,
        string memory _fromAddress,
        string memory _toAddress
    ) public returns (uint256) {
        lastOrderNumber++; // Increment the last order number to create a unique order number

        // Create a new Order struct and add it to the mapping
        orders[lastOrderNumber] = Order({
            orderNumber: lastOrderNumber,
            name: _name,
            email: _email,
            phoneNumber: _phoneNumber,
            fromAddress: _fromAddress,
            toAddress: _toAddress,
            isShipped: false // Default value at creation
        });

        // Emit the OrderCreated event
        emit OrderCreated(lastOrderNumber);

        // Return the new order number
        return lastOrderNumber;
    }

    // Function to retrieve order details by order number
    function getOrder(uint256 _orderNumber) public view returns (Order memory) {
        require(orders[_orderNumber].orderNumber != 0, "Order does not exist.");
        return orders[_orderNumber];
    }

    // Function to mark an order as shipped
    function markAsShipped(uint256 _orderNumber) public {
        require(orders[_orderNumber].orderNumber != 0, "Order does not exist.");
        orders[_orderNumber].isShipped = true;
    }
}

