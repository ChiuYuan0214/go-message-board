# TypeScript Frontend Code Style

Applies to `frontend/`. Framework: Next.js 14. Patterns extracted from the existing codebase.

---

## Components

```tsx
// 1. Union type aliases above the component
export type InputType = "text" | "email" | "password" | "phone" | "file";

// 2. Props interface explicitly typed
interface Props {
  type: InputType;
  title: string | React.ReactElement;
  value: string;
  onChange?: (val: string) => void;
  placeholder?: string;
}

// 3. Arrow function, React.FC<Props>, destructure in signature
const Input: React.FC<Props> = ({ type, title, value, onChange, placeholder }) => {
  return (
    <>
      <input
        type={type}
        value={value}
        onChange={(e) => onChange && onChange(e.target.value)}
      />
      {/* 4. styled-jsx for scoped CSS */}
      <style jsx>{`
        input {
          display: ${type === "file" ? "none" : "block"};
          width: 250px;
        }
      `}</style>
    </>
  );
};

export default Input;
```

**Rules:**
- `interface` for Props and object shapes
- `type` for unions and aliases
- Never inline props type — always a named `interface Props`
- Destructure props in the function signature, not inside the body

---

## Hooks

```tsx
// Multiple useState — one per field, not a single object
const [title, setTitle] = useState("");
const [content, setContent] = useState("");
const [tags, setTags] = useState<string[]>([]);

// useCallback for all event handlers
const handleSubmit = useCallback(async () => {
  const { status } = await api.post(GENERAL_IP + "/article", { title, content });
  if (status !== "success") return;
  router.push("/");
}, [title, content]);

// useEffect — always explicit dependency array
useEffect(() => {
  fetchArticle(articleId);
}, [articleId]);
```

---

## API Layer

All API calls go through `api/utils.ts`. Files in `api/` are grouped by resource:

```
frontend/api/
├── utils.ts        ← base fetch wrapper (get/post/put/delete)
├── auth.ts         ← /register, /login
├── article.ts      ← /article
├── articles.ts     ← /articles
├── vote.ts         ← /vote
├── profile.ts      ← /profile
└── ...
```

**Pattern for a new API function:**

```ts
// frontend/api/vote.ts
import { GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const vote = async (body: {
  sourceId: number;
  score: number;
  voteType: string;
}) => {
  return await api.post(`${GENERAL_IP}/vote`, body);
};
```

**Consuming in a component:**

```tsx
import { vote } from "@/api/vote";

const handleVote = useCallback(async (score: number) => {
  const { status } = await vote({ sourceId: articleId, score, voteType: "article" });
  if (status !== "success") return;   // ← always check status field
  // update local state
}, [articleId]);
```

**`api` object methods:** `api.get(url)`, `api.post(url, body)`, `api.put(url, body)`, `api.delete(url, body)`, `api.postForm(url, formData)`

---

## Types

```ts
// frontend/types/article.ts — interface for API response objects
export interface Article {
  articleId: number;
  userId: number;
  author: string;
  title: string;
  content: string;
  voteUp: number;
  voteDown: number;
  myScore: number;
  hasCollec: boolean;
  publishTime: string;
  tags: string[];
}
```

- `interface` for all API response shapes and context types
- `type` for unions (`type SortMode = "newest" | "hot" | "view"`)
- Field names match the JSON keys from the backend (camelCase)

---

## Imports

Use `@/` aliases:

```ts
import { Article } from "@/types/article";
import { vote } from "@/api/vote";
import { useAuth } from "@/context/auth";
```

Available aliases: `@/api`, `@/types`, `@/models`, `@/context`, `@/utils`, `@/constants`

---

## Styling

All styles use styled-jsx inside the component. No separate CSS files, no CSS modules.

```tsx
<style jsx>{`
  .container {
    display: flex;
    gap: 1rem;
  }
  .title {
    color: ${isActive ? "#333" : "#999"};  /* template literals for dynamic values */
    font-size: 1.2rem;
  }
`}</style>
```
