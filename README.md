# ChatGPT Tokenizer

This is a Go program that performs natural language processing tasks using the OpenAI API, including sentiment analysis, word frequency counting, part-of-speech tagging, and translation. It receives input text through an HTTP endpoint `/api/tokenizer` and returns the processing results in specified JSON format.

## Configuration

The program requires a valid OpenAI API key to work correctly. Please store your key in the environment variable.

- Linux/macOS:

    ```
    export OPENAI_API_KEY=your_api_key_here
    ```

- Windows:

    ```
    setx OPENAI_API_KEY "your_api_key_here"
    ```

## Docker Deployment

### Build

1. Clone or download this project to your local.
    ```
   git clone https://github.com/gopkg-dev/chatgpt-tokenizer.git
    ```

2. Run the following command in the project root directory to build a Docker image:

    ```
    docker build -t chatgpt-tokenizer .
    ```

### Run

#### With `docker run`

1. Run the following command to start a container:

    ```
    docker run -p 3000:3000 -e OPENAI_API_KEY=your_api_key_here chatgpt-tokenizer
    ```

   Replace `your_api_key_here` with your OpenAI API key.

#### With `docker-compose`

1. Save the following content to a file named `docker-compose.yml`:

    ```
    version: '3'
    services:
      chatgpt-tokenizer:
        image: chatgpt-tokenizer
        ports:
          - "3000:3000"
        environment:
          OPENAI_API_KEY: your_api_key_here
    ```

2. Run the following command to start a container:

    ```
    docker-compose up -d
    ```

   Replace `your_api_key_here` with your OpenAI API key.

## Local Deployment (non-Docker)

1. Ensure you have installed Go 1.16 or later.
2. Clone or download this project to your local.
3. Run the following command in the project root directory to build and run the program:

    ```
    go build -o chatgpt-tokenizer .
    ./chatgpt-tokenizer
    ```

## Usage

### Sending Requests

Send a POST request to the `/api/tokenizer` endpoint with the `input_text` parameter, which contains the text content to be processed.

### Response Format

The program returns the processing results in the following JSON format:

```
{
    "text":"The original text",
    "translate_cn":"Translated text",
    "language":"English",
    "words":[
        {
            "word":"Word1",
            "count":1,
            "translation_cn":"Translation1"
        },
        {
            "word":"Word2",
            "count":2,
            "translation_cn":"Translation2"
        }
    ],
    "sentiment":"Sentiment analysis result",
    "aspects":[
        {
            "aspect":"Aspect1",
            "polarity":"Polarity1",
            "translation_cn":"Translation1 of Aspect1"
        },
        {
            "aspect":"Aspect2",
            "polarity":"Polarity2",
            "translation_cn":"Translation of Aspect2"
        }
    ]
}
```

Where:

- `text`: The original input text.
- `translate_cn`: The translated text of the original text, if any translation is performed. Otherwise, it is an empty string.
- `language`: The language type of the original text.
- `words`: An array that includes detailed information about each word, including its count, translation, etc.
- `sentiment`: The sentiment analysis result of this text.
- `aspects`: An array that includes detailed information about aspects, polarity, translations, etc. related to the text.

## References

- [OpenAI API Documentation](https://beta.openai.com/docs/api-reference/introduction)
- [Go Official Documentation](https://golang.org/doc/)
- [Docker Image Official Documentation](https://docs.docker.com/engine/reference/builder/)