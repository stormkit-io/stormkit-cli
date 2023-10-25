import http from "http";

export default (req: http.IncomingMessage, res: http.ServerResponse) => {
  setTimeout(() => {
    res.setHeader("Content-Type", "text/html");
    res.end("Hello world!");
  }, 250);
};
