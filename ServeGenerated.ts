import { handleRequest } from "./StaticServer.ts";
import { parseArgs } from "@std/cli/parse-args";

const flags = parseArgs(
  Deno.args,
  {
    string: ["statics"]
  }
)

const statics = flags.statics

if (!statics) {
  throw "No generated directory configured (--statics).";
}

export default {
  fetch(req: Request) {
    return handleRequest(statics)(req)
  }
} satisfies Deno.ServeDefaultExport;
