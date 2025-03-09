# Order-Management-System

## Order Service
- Validate order details via talking with Stock Service
- CURD of Orders
- Initiates the Payment Flow by sending an event

## Stock Service
- Handles stock
- Validates order quantities

## Menu Service
-Stores items as menu

## Payment Service
- Initiates a payment with Stripe
- Produces an order Paid/Cancelled event to (orders, stock and kitchen) services
  
## Kitchen Service
- Long running process of a 'Simulated Kitchen Staff'