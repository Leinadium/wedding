<script lang="ts">
  import { onMount } from "svelte";
  import { api, type InviteResponse } from "../api";
  import Attendee from "./Attendee.svelte";
  import { loadStoredInvite, saveStoredInvite } from "./state";

  let inviteCode: string = $state("");
  let isLoading: boolean = $state(false);

  let invite: InviteResponse | undefined = $state(undefined);
  let currentNote: string = $state("");

  onMount(() => {
    let invite = loadStoredInvite();
    if (invite) {
      inviteCode = invite;
    }
  });

  $effect(() => {
    if (inviteCode.length != 6) {
      invite = undefined;
      return;
    }
    isLoading = true;
    api
      .getInvite(inviteCode)
      .then((data) => {
        invite = data;
        currentNote = invite.note;
        saveStoredInvite(inviteCode);
        isLoading = false;
      })
      .catch((e) => {
        invite = undefined;
        isLoading = false;
      });
  });

  function updateAttendeeStatusFactory(i: number) {
    return (status: boolean | null) => {
      invite!.attendees[i].confirmed = status;
    };
  }

  function updateAttendeeIsChildFactory(i: number) {
    return (isChild: boolean) => {
      invite!.attendees[i].isChild = isChild;
    };
  }

  async function saveInvite() {
    // if note is different, save it
    if (currentNote !== invite?.note) {
      invite!.note = currentNote;
      await api.saveInviteNote(inviteCode, currentNote);
    }
    // for each attendee, save
    for (let i = 0; i < invite!.attendees.length; i++) {
      const attendee = invite!.attendees[i];
      await api.saveAttendee(attendee);
    }
  }
</script>

<div class="invite">
  <div class="invite-input">
    <span>Input the code received in the invite</span>
    <input type="text" placeholder="e.g. ABC123" bind:value={inviteCode} />
  </div>
  {#if invite}
    <div class="invite-content">
      {#each invite.attendees as attendee, i}
        <Attendee
          {attendee}
          updateStatus={updateAttendeeStatusFactory(i)}
          updateIsChild={updateAttendeeIsChildFactory(i)}
        />
      {/each}
    </div>
    <div class="invite-note">
      <input
        type="text"
        placeholder="Any observations or comments"
        bind:value={currentNote}
      />
    </div>

    <div class="invite-confirm">
      <input type="submit" value="Save" onclick={saveInvite} />
    </div>
  {/if}
</div>

<style>
  .invite {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    padding: 1.5rem;
    background-color: #f9fafb;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    font-family:
      system-ui,
      -apple-system,
      sans-serif;
  }

  /* Stacks the span above the input naturally */
  .invite-input {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .invite-input span {
    font-size: 0.875rem;
    color: #4b5563;
    font-weight: 500;
  }

  /* Shared styling for text inputs */
  .invite-input input,
  .invite-note input {
    padding: 0.75rem;
    border: 1px solid #d1d5db;
    border-radius: 6px;
    font-size: 1rem;
    width: 100%;
    box-sizing: border-box;
    transition:
      border-color 0.2s,
      box-shadow 0.2s;
  }

  .invite-input input:focus,
  .invite-note input:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
  }

  .invite-content {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    padding: 1rem 0;
    border-top: 1px solid #e5e7eb;
    border-bottom: 1px solid #e5e7eb;
  }

  .invite-confirm {
    display: flex;
    justify-content: flex-end; /* Aligns the save button to the right */
  }

  .invite-confirm input[type="submit"] {
    padding: 0.75rem 2rem;
    background-color: #111827;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .invite-confirm input[type="submit"]:hover {
    background-color: #374151;
  }
</style>
