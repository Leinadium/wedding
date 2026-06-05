<script lang="ts">
  import type { Attendee } from "../api";

  let {
    attendee,
    updateStatus,
    updateIsChild,
  }: {
    attendee: Attendee;
    updateStatus: (status: boolean | null) => void;
    updateIsChild: (status: boolean) => void;
  } = $props();

  function statusToText(status: boolean | null): string {
    switch (status) {
      case true:
        return "Confirmed!";
      case false:
        return "Won't attend.";
      default:
        return "Pending...";
    }
  }

  function statusToClass(status: boolean | null): string {
    switch (status) {
      case true:
        return "yes";
      case false:
        return "no";
      default:
        return "pending";
    }
  }

  let status: boolean | null = $derived(attendee.confirmed);
  let statusClass: string = $derived(statusToClass(status));
  let statusText: string = $derived(statusToText(status));

  function toggleStatus() {
    switch (status) {
      case false:
        updateStatus(true);
        break;
      case true:
        updateStatus(null);
        break;
      default:
        updateStatus(false);
        break;
    }
  }

  function toggleIsChild() {
    updateIsChild(!attendee.isChild);
  }
</script>

<div class="attendee">
  <button class="child" onclick={toggleIsChild}>
    {attendee.isChild ? "Child" : "Adult"}
  </button>

  <span class="name">{attendee.name}</span>

  <button class="status {statusClass}" onclick={toggleStatus}>
    {statusText}
  </button>
</div>

<style>
  .attendee {
    display: flex;
    flex-flow: row nowrap;
    justify-content: right;
    align-items: center;
    gap: 1rem; /* Slightly wider gap for better spacing */
    border-radius: 6px;

    font-family: cursive;
    font-size: 1.3rem;
  }

  .name {
    flex: 1;
    width: 100%;
    font-weight: 500;
    color: #1f2937;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;
  }

  button {
    background-color: rgba(0, 0, 0, 0.05);

    cursor: pointer;
    padding: 0.2rem 0.3rem;
    font-weight: 500;
    font-size: 1.2rem;
    font-family: cursive;
    text-decoration: underline;

    border: 1px solid transparent;
    border-radius: 5px;

    transition: border 0.2s ease-in-out;
  }

  button:hover {
    border: 1px solid #767;
  }

  .child {
    color: #767;
  }

  :global(.status.yes) {
    color: #166534;
    background-color: #16653430;
  }

  :global(.status.no) {
    color: #991b1b;
    background-color: #991b1b30;
  }

  :global(.status.pending) {
    color: #92400e;
    background-color: #99999930;
  }
</style>
