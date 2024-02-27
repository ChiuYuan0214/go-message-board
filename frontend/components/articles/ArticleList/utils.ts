import {
  getArticles,
  getHotArticles,
  getProfileArticles,
  getViewArticles,
} from "@/api/articles";
import { getCollections } from "@/api/collections";

export const getData = async (
  type: string,
  userId: number,
  page: number,
  isSSR: boolean
) => {
  switch (type) {
    case "hot":
      return await getHotArticles(userId, page, isSSR);
    case "view":
      return await getViewArticles(userId, page, isSSR);
    case "profile":
      return await getProfileArticles(userId, page);
    case "collection":
      return await getCollections(page);
    default:
      return await getArticles(userId, page, isSSR);
  }
};
