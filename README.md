# Events Aggregate

## STEPS

1. **Install Go on your system:**

   Follow the official [Go installation instructions](https://golang.org/doc/install) to install Go on your machine.

2. **Clone this repository:**

   ```bash
   git clone https://github.com/rebel532/events_aggregate.git

   testcase1 : go run aggregate_events.go -i <input_file_path> -o <output_file_path>
   exampple :  go run aggregate_events.go -i ~/Desktop/events_aggregate/src/core/input.json -o ~/Desktop/events_aggregate/src/core/output.json

   testcase2 : go run aggregate_events.go -i <input_file_path> -o <output_file_path> --update
   example: go run aggregate_events.go -i ~/Desktop/events_aggregate/src/core/input.json -o ~/Desktop/events_aggregate/src/core/output.json --update

3. output testcse : 1
   ```json
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

4. output testcase : 2
    ```json
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

5. Observation : second time counts increases for each user Ids

