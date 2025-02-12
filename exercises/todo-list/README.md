   ## Instruction
   The Story of the To-Do List Manager

Every day, Mia had many things to do—homework, meetings, shopping, and more. But there was a problem. She kept forgetting important tasks because she wrote them on random papers or tried to remember them in her head.

*"I need a better way to keep track of my tasks!"* she thought.

So, Mia decided to create **The To-Do List Manager**—a simple tool to help her stay organized. Her system would allow her to:

✅ **Add new tasks** whenever she needed.

✅ **Remove tasks** once they were done.

✅ **See all her tasks** in one clear list.

With a plan in mind, she started building her perfect to-do list. Would this system help Mia stay organized and never forget a task again? That was up to her…

   ## Flow building To-Do List Manager
 
   1. Problem Analysis
      - Hard to track tasks

   2. Decomposition
      Components:
      - Input task
      - Task storage [name task and status task]
      - See all tasks
      - Remove function when task is done

   3. Pattern Recognition 
      - CRUD operation (Create, Read, and Delete)
      - Input/Output
      - Data persistence
   
   4. Essential Focus
      Must Have:
      - Task name
      - Remove task capability
      - Data persistence

      Can Skip:
      - Complex UI
      - Advanced features

   5. Algorithm Flow
      - Start program
      - Show options 
      - Process show
      - [option] Create new task
         - Input task name
         - Save task to task storage
         - Show all task
         - End program
      - [option] Complete / remove task
         - Input task name
         - Say it's Complete
         - Remove task
         - End program
      - [option] list task
         - Show all task
         - End program
      - [option] Exit program