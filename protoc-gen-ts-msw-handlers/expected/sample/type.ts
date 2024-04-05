export type Tag = {
  id: string;
  name: string;
  color: string;
};

export type GetTagsResponse = {
  tags: Tag[];
};

export type GetTagsRequest = {
  limit: number;
  smallestTagId: string;
  biggestTagId: string;
  w: string;
};
