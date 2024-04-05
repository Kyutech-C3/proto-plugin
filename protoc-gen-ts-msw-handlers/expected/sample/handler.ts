import { http, HttpResponse } from "msw";
export const V1TagsGetHandler = http.get("/v1/tags", () => {
  return new HttpResponse.json([{ id: "1", name: "tag1", color: "#000000" }]);
});
