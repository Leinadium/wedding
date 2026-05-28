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
        return "✅ Will be there!";
      case false:
        return "❌ Won't go :(";
      default:
        return "⏳ Pending";
    }
  }

  function statusToClass(status: boolean | null): string {
    switch (status) {
      case true:
        return "confirmed";
      case false:
        return "not-confirmed";
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
  <span class="name">{attendee.name}</span>

  <button class="status {statusClass}" onclick={toggleStatus}>
    {statusText}
  </button>

  <div class="child-container">
    <span class="child-label">Child?</span>
    <input
      type="checkbox"
      checked={attendee.isChild}
      onchange={toggleIsChild}
    />
  </div>
</div>

<style>
  .attendee {
    display: flex;
    align-items: center;
    gap: 1.5rem; /* Slightly wider gap for better spacing */
    padding: 0.75rem 1rem;
    background-color: white;
    border: 1px solid #e5e7eb;
    border-radius: 6px;
  }

  .name {
    flex: 1; /* Pushes everything else to the right */
    font-weight: 500;
    color: #1f2937;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .status {
    /* FIX: Fixed width and flex-shrink prevents layout shifting */
    width: 120px;
    flex-shrink: 0;
    text-align: center;

    padding: 0.5rem 0; /* Removed horizontal padding to let width control it */
    border: 1px solid #d1d5db;
    border-radius: 6px;
    background-color: white;
    color: #374151;
    font-size: 0.875rem;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .status:hover {
    background-color: #f3f4f6;
  }

  :global(.status.confirmed) {
    background-color: #dcfce7;
    color: #166534;
    border-color: #bbf7d0;
  }

  :global(.status.declined) {
    background-color: #fee2e2;
    color: #991b1b;
    border-color: #fecaca;
  }

  /* NEW: Stacks the label on top of the checkbox */
  .child-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.25rem;
    flex-shrink: 0; /* Prevents the checkbox area from squeezing */
  }

  .child-label {
    font-size: 0.75rem;
    color: #6b7280;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.05em;
  }

  input[type="checkbox"] {
    width: 1.25rem;
    height: 1.25rem;
    cursor: pointer;
    accent-color: #111827;
    margin: 0;
  }
</style>
