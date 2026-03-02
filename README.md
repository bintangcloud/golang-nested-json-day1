# Day 1 - Golang Backend 🚀  
## Nested JSON: Order, Customer, and Items

Today I learned how to create nested JSON structures in Golang using struct and slice.

---

## 📌 Problem

Create a struct `Order` that contains:
- A nested struct `Customer`
- A slice (list) of `Item`
- Convert the struct into JSON format

---

## 🏗️ Data Structure Design

The structure looks like this:

Order
- ID
- Customer
  - Name
  - Email
- Items (list)
  - Name
  - Price
  - Quantity

---
