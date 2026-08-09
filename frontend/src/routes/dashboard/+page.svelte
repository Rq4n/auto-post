<script lang="ts">
  import { createPost } from "./post";

  async function handleCreatePost() {
    try {
      const data = await createPost({
        title: "teste",
        content: "teste",
      });

      console.log("Post criado:", data);
    } catch (error) {
      console.error(error);
    }
  }
</script>

<div class="flex h-screen">
  <aside class=" flex flex-col w-64 border-r border-zinc-800">
    <!-- logo -->
    <div class="flex px-6 py-4 justify-start">
      <a href="/"> <span class="text-blue-600">/</span>autopost</a>
    </div>
    <hr />

    <!-- Create post CTA -->
    <div class="flex py-6 px-4 justify-start">
      <button
        onclick={handleCreatePost}
        class=" text-border bg-blue-600 text-white hover:bg-blue-700 shadow-lg text-md px-4"
        >Create Post</button
      >
    </div>

    <!-- Subtopic posts -->
    <div class="flex flex-col py-4 m-6">
      <p class="mb-2 text-xs uppercase text-zinc-500">Content</p>
      <div class="flex flex-col underline gap-2">
        <a href="/posts" class="hover:bg-zinc-200">Posts</a>
        <a href="/calendar" class="hover:bg-zinc-200">Calendar</a>
        <a href="/schedule" class="hover:bg-zinc-200">Schedule</a>
      </div>
    </div>

    <!-- Subtopic configs -->
    <div class="flex flex-col m-6 py-4">
      <p class="mb-2 text-xs uppercase text-zinc-500">Configuration</p>
      <div class="flex flex-col underline gap-2">
        <a href="/connections" class="hover:bg-zinc-200">Connections</a>
        <a href="/settings" class="hover:bg-zinc-200">Settings</a>
      </div>
    </div>

    <!-- Footer -->
    <div class="mt-auto pb-6">
      <hr />
      <div class="flex flex-col gap-6 py-4 px-5 pr-20">
        <button
          class="flex px-4 bg-blue-600 hover:bg-blue-700 shadow-lg text-white"
          >Upgrade</button
        >
        <button class="flex px-4 hover:bg-zinc-200">Logout</button>
      </div>
    </div>
  </aside>

  <main class="flex-1 p-8">Conteúdo da página</main>
</div>
