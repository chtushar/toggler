const ivm = require("isolated-vm");

async function main() {
  let isolate = new ivm.Isolate({ memoryLimit: 20 }); // Memory limit in MB
  let context = await isolate.createContext(); // Create a persistent context

  const executeCode = async (code) => {
    try {
      // Create a script
      const script = await isolate.compileScript(code);

      // Run the script in the context
      await script.run(context);

      const global = context.global;
      const result = await global.get("result");

      return await result;
    } catch (error) {
      return `Error: ${error.message}`;
    }
  };

  const processCode = (data) => {
    executeCode(data.toString())
      .then((result) => {
        process.stdout.write(JSON.stringify(result) + "\n");
      })
      .catch((error) => {
        console.error(`Execution error: ${error.message}`);
      });
  };

  // Read from stdin
  process.stdin.on("data", processCode);
}

main().catch((error) =>
  console.error(`Failed to initialize: ${error.message}`)
);
