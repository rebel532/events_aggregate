# Event Aggregate
## Instructions

1. Install Go on your system.

2. Clone this repository.

3. in cmd folder run :
    for Testcase 1 :  % go run aggregate_events.go -i <path of input json> -o <path of output json>
    for Testcase 2 :  % go run aggregate_events.go -i <path of input json> -o <path of output json>  --update
   
    Example : 
    Test case 1 : rahul@FIN-IT-LT-524-MAC cmd % go run main.go -i ~/Desktop/events_aggregate/src/core/input.json -o ~/Desktop/events_aggregate/src/core/output.json
    Test case 2 : rahul@FIN-IT-LT-524-MAC cmd % go run main.go -i ~/Desktop/events_aggregate/src/core/input.json -o ~/Desktop/events_aggregate/src/core/output.json --update

4. For test case 1: run just run first command
    output.json will look like : 
    [
        {
            "date": "2022-12-31",
            "likeReceived": 2,
            "post": 1,
            "userId": 1
        },
        {
            "date": "2023-01-01",
            "post": 1,
            "userId": 1
        },
        {
            "comment": 1,
            "date": "2023-01-01",
            "post": 1,
            "userId": 2
        }
    ]
5. For test case 2: run the send command
   output.json will look like : 
    [
        {
            "date": "2022-12-31",
            "likeReceived": 4,
            "post": 2,
            "userId": 1
        },
        {
            "date": "2023-01-01",
            "post": 2,
            "userId": 1
        },
        {
            "comment": 2,
            "date": "2023-01-01",
            "post": 2,
            "userId": 2
        }
    ]

6. Observation : 
   second time counts increases for  each user Ids
