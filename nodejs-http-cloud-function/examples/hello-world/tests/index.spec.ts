import { Request, Response } from "express";
import { It, Mock, Times } from "typemoq";

import { helloWorld } from "../src";

describe("index.ts", () => {
  describe("helloWorld(request, response)", () => {
    it("sets a status code of 200", () => {
      // Arrange
      const mockRequest = Mock.ofType<Request>();
      const mockResponse = Mock.ofType<Response>();

      mockResponse
        .setup(r => r.status(It.isAny()))
        .returns(() => mockResponse.object);

      // Act
      helloWorld(mockRequest.object, mockResponse.object);

      // Assert
      mockResponse.verify(r => r.status(200), Times.once());
    });

    it("sends a message containing 'Hello world!'", () => {
      // Arrange
      const mockRequest = Mock.ofType<Request>();
      const mockResponse = Mock.ofType<Response>();

      mockResponse
        .setup(r => r.status(It.isAny()))
        .returns(() => mockResponse.object);

      // Act
      helloWorld(mockRequest.object, mockResponse.object);

      // Assert
      mockResponse.verify(
        r => r.send(It.is<string>(b => b.includes("Hello world!"))),
        Times.once()
      );
    });
  });
});
