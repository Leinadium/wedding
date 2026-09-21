<script lang="ts">
  import type { Attendee } from "../../api";
  import { timeAgo } from "../../common";

  let {
    at,
  }: {
    at: Attendee;
  } = $props();

  async function copyToClipboard(text: string) {
    await navigator.clipboard.writeText(text);
  }
</script>

<tr>
  <td class="u-truncate">{at.name}</td>
  <td>
    {#if at.confirmed}
      <span class="p-chip--positive is-readonly is-inline is-dense">
        Confirmed
      </span>
    {:else if at.confirmed === false}
      <span class="p-chip--negative is-readonly is-inline is-dense">
        Won't go
      </span>
    {:else}
      <span class="p-chip--caution is-readonly is-inline is-dense">
        To be confirmed
      </span>
    {/if}
  </td>

  <td>{timeAgo(at.updatedAt)}</td>

  <td class="uuid-cell">
    <div class="uuid-container">
      <code class="u-truncate" title={at.inviteId}>
        {at.inviteId}
      </code>
      <button
        class="copy-btn"
        aria-label="Copy Invite ID to clipboard"
        onclick={() => copyToClipboard(at.inviteId)}
      >
        <i class="p-icon--copy"></i>
      </button>
    </div>
  </td>
</tr>

<style>
  .uuid-cell {
    max-width: 100px;
  }

  .uuid-container {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 0.5rem;
  }

  .copy-btn {
    background: transparent;
    border: none;
    padding: 0;
    margin: 0;
    cursor: pointer;
    display: flex;
    align-items: center;
  }

  .copy-btn:hover {
    opacity: 0.7;
  }
</style>
