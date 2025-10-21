import { handleRequest } from "./StaticServer.ts";
import { env } from "node:process";

const statics = env.GENERATED;

if (!statics) {
  throw "No generated directory configured (env GENERATED).";
}

export default {
  fetch(req: Request) {
    return handleRequest(statics)(req)
  }
} satisfies Deno.ServeDefaultExport;
