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
  <input type="checkbox" checked={attendee.isChild} onchange={toggleIsChild} />
</div>
