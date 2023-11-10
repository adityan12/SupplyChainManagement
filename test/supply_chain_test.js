const SupplyChain = artifacts.require("SupplyChain");

contract("SupplyChain", accounts => {
  let supplyChain;
  const sender = accounts[0];
  const receiver = accounts[1];
  
  before(async () => {
    supplyChain = await SupplyChain.deployed();
  });

  describe("sendOrder", () => {
    it("should create a new order", async () => {
      const result = await supplyChain.createOrder("Alice", "alice@example.com", "1234567890", "123 Sender Address", "456 Receiver Address", { from: sender });
      assert.equal(result.logs[0].event, "OrderCreated", "OrderCreated event should be emitted");
    });
  });

  describe("getOrder", () => {
    it("should retrieve an order by orderNumber", async () => {
      const orderNumber = 1; // Assuming the first order has orderNumber 1
      const order = await supplyChain.getOrder(orderNumber, { from: sender });
      
      assert.equal(order.name, "Alice", "The name of the retrieved order should be Alice");
      // Add more assertions as needed for other order details
    });
  });

  describe("markAsShipped", () => {
    it("should mark an order as shipped", async () => {
      const orderNumber = 1; // The orderNumber to mark as shipped
      await supplyChain.markAsShipped(orderNumber, { from: sender });

      const order = await supplyChain.getOrder(orderNumber, { from: sender });
      assert.equal(order.isShipped, true, "The order should be marked as shipped");
    });
  });
});

