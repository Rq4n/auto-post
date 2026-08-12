<script lang="ts">
  import { LoginWithGoogle } from "$lib/auth/auth";
</script>

<div>
  <!-- navbar -->
  <div class="flex justify-between font-bold items-center mx-40 my-4 text-sm">
    <a href="/"><span class="text-blue-600">/</span>autopost</a>

    <div class="flex items-center gap-4 font-bold text-xs">
      <a href="/pricing" class="hover:underline">Pricing</a>
      <button onclick={LoginWithGoogle} class="text-gray-600 hover:text-black">
        Login
      </button>
    </div>
  </div>

  <section class="mx-40 my-20 flex flex-col items-center text-center">
    <!-- headline -->
    <h1 class="text-4xl font-extrabold tracking-tight max-w-xl leading-tight">
      Publish everywhere on <span
        class="bg-blue-100 text-blue-700 px-2 py-0.5 rounded">autopilot</span
      >
    </h1>

    <!-- subheadline -->
    <p class="text-gray-600 text-base mt-3 max-w-md font-medium">
      Schedule once. Reach your audience on all platforms automatically.
    </p>

    <!-- cta -->
    <div class="mt-6 flex flex-col items-center gap-2">
      <button
        onclick={LoginWithGoogle}
        class="bg-blue-700 hover:bg-blue-800 text-white font-bold py-3 px-6 text-sm transition-all shadow-md rounded"
      >
        Start auto-posting now →
      </button>

      <span class="text-xs text-gray-500 font-medium">
        Free forever • No credit card needed
      </span>
    </div>

    <!-- demo placeholder -->
    <div
      class="mt-10 w-full max-w-3xl border border-gray-200 bg-gray-50 rounded-lg p-4 text-xs text-gray-400"
    >
      [ Add your product screenshot or quick demo video here ]
    </div>
  </section>
</div>
