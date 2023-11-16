Why: PORT?
- As we tend to use Clean Architecture for separating the business & other things.
- We have here two main modules:
    1. High level modules: which contain complex logic, should be easily reusable and unaffected by changes in low-level modules, which provide utility features.
    2. Low level modules: are bunch of small modules that help high-level modules to do their works.
- Normally in N-layer architecture, we have presentation layer -> business layer -> data access layer. That would not be a good idea, so with clean architect, we have: present -> business <- data access layer. (actually, present + data access = infrastructure layer). By using port, applied Depedency Inversion Principle, we inverse the dependencies direction, keep the business layers high-level.