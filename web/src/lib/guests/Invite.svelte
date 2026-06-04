<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import { onMount } from "svelte";
  import { api, type InviteResponse } from "../api";
  import Attendee from "./Attendee.svelte";
  import { loadStoredInvite, saveStoredInvite } from "./state";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

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

<div class="invite-wrapper" transition:fade={{ duration: 300 }}>
  <div class="invite" transition:fly={{ duration: 300, y: +150 }}>
    <div class="input">
      <span>Input the code received in the invite</span>
      <input type="text" placeholder="ABC123" bind:value={inviteCode} />
    </div>
    {#if invite}
      <div class="content">
        {#each invite.attendees as attendee, i}
          <Attendee
            {attendee}
            updateStatus={updateAttendeeStatusFactory(i)}
            updateIsChild={updateAttendeeIsChildFactory(i)}
          />
        {/each}
      </div>
      <div class="note">
        <input
          type="text"
          placeholder="Any observations or comments"
          bind:value={currentNote}
        />
      </div>
      <div class="confirm">
        <input type="submit" value="Save" onclick={saveInvite} />
      </div>
    {/if}
    <button class="close" onclick={closeCb}>X</button>
  </div>
</div>

<style>
  .invite-wrapper {
    position: fixed;
    top: 0;
    left: 0;

    width: 100vw;
    height: 100vh;

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    background-color: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }

  .invite {
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    padding: 2rem;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    font-family:
      system-ui,
      -apple-system,
      sans-serif;

    background-image: url("src/assets/images/texture.png");
    background-color: #f0f0f0;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    width: 80%;
    max-width: 800px;
    min-height: 500px;
    height: auto;
    max-height: 90vh;
    overflow-y: auto;
  }

  /* Stacks the span above the input naturally */
  .input {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;

    display: flex;
    flex-flow: column nowrap;
    align-items: center;
    justify-content: center;
  }

  .input span {
    font-size: 0.875rem;
    color: #4b5563;
    font-weight: 500;
  }

  /* Shared styling for text inputs */
  .input input,
  .note input {
    padding: 0.75rem;
    border: 1px solid #d1d5db;
    border-radius: 6px;
    font-size: 1rem;
    width: 100%;
    box-sizing: border-box;
    text-align: center;
    transition:
      border-color 0.2s,
      box-shadow 0.2s;

    width: 80%;
    max-width: 100px;
  }

  .input input:focus,
  .note input:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
  }

  .content {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    padding: 1rem 0;
    border-top: 1px solid #e5e7eb;
    border-bottom: 1px solid #e5e7eb;
  }

  .confirm {
    display: flex;
    justify-content: flex-end; /* Aligns the save button to the right */
  }

  .confirm input[type="submit"] {
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

  .confirm input[type="submit"]:hover {
    background-color: #374151;
  }

  .close {
    position: absolute;
    top: 1rem;
    right: 1rem;

    background: none;
    border: none;
    font-size: 1.5rem;
    color: #777;

    cursor: pointer;
  }
</style>
