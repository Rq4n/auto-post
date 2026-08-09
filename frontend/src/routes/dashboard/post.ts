interface createPostBody {
  title: string;
  content: string;
}

export async function createPost(body: createPostBody) {
  const res = await fetch("http://localhost:8080/api/posts", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(body),
  });

  if (!res.ok) {
    const message = await res.text();
    throw new Error(message);
  }
  return await res.json();
}
