import { Request, Response } from "express";

const helloWorld = (request: Request, response: Response) => {
  response.status(200).send("Hello world!");
};

export { helloWorld };
