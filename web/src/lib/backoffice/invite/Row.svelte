<script lang="ts">
  import type { Invite } from "../../api";
  let {
    invite,
  }: {
    invite: Invite;
  } = $props();

  const names: string = $derived(invite.attendees.map((a) => a.name).join(","));

  let expanded: boolean = $state(false);

  async function copyToClipboard(text: string) {
    await navigator.clipboard.writeText(text);
  }
</script>

<tr>
  <td class="uuid-cell">
    <div class="uuid-container">
      <code title={invite.id}>
        {invite.id}
      </code>
      <button
        class="copy-btn"
        aria-label="Copy ID to clipboard"
        onclick={() => copyToClipboard(invite.id)}
      >
        <i class="p-icon--copy"></i>
      </button>
    </div>
  </td>

  <td class="uuid-cell">
    <div class="uuid-container">
      <code title={invite.phone}>
        {invite.phone}
      </code>
      <button
        class="copy-btn"
        aria-label="Copy ID to clipboard"
        onclick={() => copyToClipboard(invite.phone)}
      >
        <i class="p-icon--copy"></i>
      </button>
    </div>
  </td>

  <td class="u-truncate">
    <span class="p-chip--information is-readonly is-inline is-dense">
      {invite.attendees.length}
    </span>
    {names}
  </td>

  {#if invite.note}
    <td class="u-align--left">
      <button
        class="u-toggle is-dense"
        aria-controls="expanded-row"
        aria-expanded={expanded ? "true" : "false"}
        onclick={(e) => {
          expanded = !expanded;
        }}
      >
        Note {expanded ? "close" : ""}
      </button>
    </td>
    <td
      id="expanded-row"
      class="p-table__expanding-panel"
      aria-hidden={!expanded ? "true" : "false"}
    >
      <div class="row">
        <div class="col-8">
          <p>{invite.note}</p>
        </div>
      </div>
    </td>
  {:else}
    <td class="u-align--left">
      <span class="p-chip--negative is-readonly is-inline is-dense">
        No note
      </span>
    </td>
    <td aria-hidden="true"></td>
  {/if}
</tr>

<style>
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
