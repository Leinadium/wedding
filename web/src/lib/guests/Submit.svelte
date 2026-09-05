<script lang="ts">
  import { getText } from "../text/text";

  let {
    onClick,
    success,
  }: {
    onClick: () => void;
    success: boolean;
  } = $props();

  function handle() {
    if (!success) onClick();
  }
</script>

<input
  class={["save", "formal", { success }]}
  type="submit"
  value={getText(success ? "invite-saved" : "invite-save")}
  onclick={handle}
/>

<style>
  .save {
    padding: 0.5rem 1rem;
    background-color: transparent;
    color: #dedacd;
    border: none;
    border-radius: 6px;
    border: 3px solid #dedacd;
    font-size: 1.2rem;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .success {
    background-color: #777155;
  }
</style>
