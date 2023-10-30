import * as core from "@actions/core";
import * as app from "./app";

/**
 * The main function for the action.
 * @returns {Promise<void>} Resolves when the action is complete.
 */
export async function run(): Promise<void> {
  try {
    const a = app.load();
    core.debug(JSON.stringify(a));
    core.debug(`Yaml: ${app.toYaml(a)}`);

    // Set outputs for other workflow steps to use
    core.setOutput("time", new Date().toTimeString());
  } catch (error) {
    // Fail the workflow run if an error occurs
    if (error instanceof Error) core.setFailed(error.message);
  }
}
