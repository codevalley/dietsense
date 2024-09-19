# Workflow Documentation

## Introduction

This document outlines the workflow for developing our app. The workflow is designed to ensure efficient progress tracking, clear communication, and alignment with our goals and design principles. We maintain two primary documents for day-to-day operations and refer to several foundational documents that capture our project's essence, design language, and development plan.

---

## Overview of Documents

### Primary Documents

1. **progress.md**

   - **Purpose**: Captures the current chunk under development, including both completed tasks and tasks to be completed (in progress).
   - **Usage**: Updated regularly to reflect the latest status of tasks within the active development chunk.
   - **Content**:
     - Current development chunk title.
     - List of tasks:
       - Completed tasks with checkmarks or strikethroughs.
       - In-progress tasks with status updates.
       - Pending tasks to be started.
     - Notes on any blockers or issues encountered.
   - **References**: Links to relevant sections in `Approach.md`, `ProjectPlan.md` for context.


### Foundational Documents


1. **Approach.md**

   - **Purpose**: Details the features, user journeys, and workflows discussed during the planning phase.
   - **Content**:
     - In-depth descriptions of app functionalities.
     - User flow diagrams and narratives.
     - Decisions on feature priorities and exclusions.

2. **ProjectPlan.md**

   - **Purpose**: Outlines the development plan, broken into 10 chunks with milestones and sub-tasks.
   - **Content**:
     - Detailed breakdown of development phases.
     - Task lists with dependencies and estimated timelines.
     - Milestones with verifiable outcomes.

---

## Workflow Process

### Development Cycle

1. **Selecting the Current Development Chunk**

   - From `ProjectPlan.md`, identify the next chunk scheduled for development.
   - Update `progress.md` to reflect the current chunk, including its title and description.

2. **Task Management within the Chunk**

   - List all tasks associated with the current chunk in `progress.md`.
   - For each task:
     - Provide a brief description.
     - Assign responsible team members.
     - Set status indicators (To Do, In Progress, Completed).

3. **Updates**

   - Update `progress.md` at the end of each code update.
   - Include progress made, challenges faced, and any adjustments to timelines.
   - Use checkmarks or status labels to indicate task completion or progress.

4. **Referencing Foundational Documents**

   - When working on tasks, refer to the relevant sections in:
     - `Approach.md` for detailed feature descriptions and user journeys.
     - `ProjectPlan.md` for context on how tasks fit into the overall project timeline.

5. **Issue Tracking and Resolution**

   - Document any blockers or issues in `progress.md`.
   - Assign action items for resolving issues.
   - If the issues are carried forward to the next chunk, keep them in progress.md under issues section (with reference to the chunk & task number)

---

## Maintenance and Update Procedures

### Updating progress.md

- **Task Completion**

  - When a task is completed, mark it as such in `progress.md`.
  - Provide any relevant notes or comments on the implementation.

- **Adding New Tasks**

  - If new tasks arise within the current chunk, add them to `progress.md` with appropriate details.
  - Ensure they align with the objectives outlined in `ProjectPlan.md`.

- **In-Progress Tasks**

  - For tasks that are underway, include brief status updates.
  - Note any expected completion dates or shifts in timelines.

### Referencing Other Documents

- **Cross-Referencing**

  - In `progress.md`, include hyperlinks or references to specific sections in the foundational documents when relevant.
  - Example: "See `DesignLanguage.md` section on 'Buttons and CTAs' for guidelines."

- **Document Updates**

  - Keep `Approach.md` or `ProjectPlan.md` like an immutable document and add any suggestions to the bottom of the document under "Suggestion" section, note these in `progress.md`.
  - Summarize the changes and their impact on current tasks.

### Document Versioning

- **Backup and Archiving**

  - When a chunk is completed, compact the contents of `Progress.md` into a `Completed.md` file and flush the contents of `Progress.md`

---

## Appendices

### Sample progress.md Structure

```markdown
# Progress - [Current Chunk Title]

## Chunk Overview

- **Description**: [Brief description of the current development chunk]
- **Milestone**: [Expected outcome of the chunk]

## Tasks

### Completed Tasks

- [x] **Task 1**: Description of task 1
  - Notes: Any relevant comments or findings
- [x] **Task 2**: Description of task 2
  - Notes: Reference to `Approach.md` section on typography

### In-Progress Tasks

- [ ] **Task 3**: Description of task 3
  - Status: In Progress
  - Assigned to: Team Member Name
  - Notes: Encountered an issue with [specific detail], see `Approach.md` for context

### Pending Tasks

- [ ] **Task 4**: Description of task 4
  - Status: To Do
  - Assigned to: Team Member Name

## Issues and Blockers

- **Issue 1**: Description of the issue
  - Action Items: Steps being taken to resolve
  - References: See `ProjectPlan.md` Chunk 5 for related tasks

## Notes

- **General Observations**: Any overarching comments or insights
- **Adjustments**: Any changes to timelines or task scopes

## References

- [Approach.md](./Approach.md)
- [ProjectPlan.md](./ProjectPlan.md)
```

---

### Document Maintenance Tips

- **Consistency**: Ensure all team members follow the same format when updating documents.
- **Clarity**: Use clear and concise language to avoid misunderstandings.
- **Timeliness**: Update documents promptly to reflect the most current information.
- **Collaboration**: Encourage open communication about changes or updates made to any documents.

---

By adhering to this workflow and utilizing the outlined documents effectively, we can maintain a high level of organization and focus throughout the development process, ultimately leading to the successful completion of our app.