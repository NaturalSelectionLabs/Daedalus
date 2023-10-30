/**
 * Unit tests for the action's main functionality, src/main.ts
 *
 * These should be run as if the action was called from a workflow.
 * Specifically, the inputs listed in `action.yml` should be set as environment
 * variables following the pattern `INPUT_<INPUT_NAME>`.
 */

import * as core from "@actions/core";
import * as main from "../src/main";

// Mock the GitHub Actions core library
const debugMock = jest.spyOn(core, "debug");
const getInputMock = jest.spyOn(core, "getInput");
const setFailedMock = jest.spyOn(core, "setFailed");
const setOutputMock = jest.spyOn(core, "setOutput");

// Mock the action's main function
const runMock = jest.spyOn(main, "run");

// Other utilities
const timeRegex = /^\d{2}:\d{2}:\d{2}/;

describe("action", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("sets the time output", async () => {
    // Set the action's inputs as return values from core.getInput()
    getInputMock.mockImplementation((name: string): string => {
      switch (name) {
        case "name":
          return "daedalus";
        case "namespace":
          return "guardian";
        case "cluster":
          return "dev";
        case "project":
          return "guardian";
        case "repo":
          return "test/repo";
        case "helm-value-files":
          return "deploy/dev/values.yaml\ndeploy/dev/values.yaml";
        default:
          return "";
      }
    });

    await main.run();
    expect(runMock).toHaveReturned();

    // Verify that all of the core library functions were called correctly
    expect(setOutputMock).toHaveBeenNthCalledWith(
      1,
      "time",
      expect.stringMatching(timeRegex),
    );
  });

  //   it("sets a failed status", async () => {
  //     // Set the action's inputs as return values from core.getInput()
  //     getInputMock.mockImplementation((name: string): string => {
  //       switch (name) {
  //         case "milliseconds":
  //           return "this is not a number";
  //         default:
  //           return "";
  //       }
  //     });

  //     await main.run();
  //     expect(runMock).toHaveReturned();

  //     // Verify that all of the core library functions were called correctly
  //     expect(setFailedMock).toHaveBeenNthCalledWith(
  //       1,
  //       "milliseconds not a number"
  //     );
  //   });
});