import * as fs from "@std/fs";

export function handleRequest(statics: string) {
  return async (req: Request) => {
    if (req.method !== "GET") {
      return new Response(
        "",
        {
          status: 404,
        },
      );
    }

    const url = new URL(req.url);

    const path = url.pathname;

    const filepath = statics + path;
    const fileExists = await fs.exists(filepath);

    if (!fileExists) {
      return new Response(
        "",
        {
          status: 404,
        },
      );
    }

    const content = await Deno.readTextFile(filepath);

    let contentType = "";
    if (filepath.endsWith(".html")) {
      contentType = "text/html; charset=utf-8";
    } else if (filepath.endsWith(".css")) {
      contentType = "text/css; charset=utf-8";
    }

    return contentType
      ? new Response(
        content,
        {
          status: 200,
          headers: {
            "Content-Type": contentType,
          },
        },
      )
      : new Response(
        "",
        {
          status: 404,
        },
      );
  };
}

export function staticServer(statics: string) {
  const server = Deno.serve(handleRequest(statics));

  return server;
}
