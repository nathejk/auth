# Keeping track of asset status

## Produces the following events:

- *acquired* first time an asset is seen
- *arrived* When an asset initially arrives at the depot ready for rental
- *use.started* in catch-up phase: when a rentalOrder is seen, in live phase: when status changes to 'rented' (or 'rentedworkshop')
- *use.ended* in catch-up phase: when next rentalOrder is seen, previous rentalOrder ends, in live phase: when status changes away from 'rented' (or 'rentedworkshop')
- *sold* when status changes to 'sold', will run after catch-up of all rentalOrders.

## Event structure

### acquired

- body
    - assetId
    - ownerId
- meta
    - producer: "ax-bridge-assetstatus"
    - sourceEvents: ["channel:type:sequence", "ax:rentalorder.synced:11223344"]
    - phase: "catchup" | "live"

### use.started

- body
    - assetId
    - customerId
    - rentalOrder: struct | null
        - rentalId:"rental-f62460c5-37e5-4dd9-bd9b-842d49159fa7"
        - shortTerm: boolean
        - OnRentDateTime:"2016-06-01T07:00:00Z"
        - expectedOffRentDateTime:"2019-05-24T07:00:00Z"
        - offRentDateTime:"2019-02-25T18:00:00Z"
- meta
    - producer: "ax-bridge-assetstatus"
    - sourceEvents: ["channel:type:sequence", "ax:rentalorder.synced:11223344"]
    - phase: "catchup" | "live"


### use.ended

- body
    - assetId
    - customerId
    - rentalOrderId: string | null
- meta
    - producer: "ax-bridge-assetstatus"
    - sourceEvents: ["channel:type:sequence", "ax:rentalorder.synced:11223344"]
    - phase: "catchup" | "live"


