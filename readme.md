
### Creating a simple tag using the last commit message and the last tag.

<ul>
<li>If the commit message contains #major or [major] then the major version will be increased by 1.</li>
<li>If the commit message contains #minor or [minor] or feat then the minor version will be increased by 1.</li>
<li>If the commit message contains #patch or [patch], or none of the above, it will be incremented by 1.</li>
</ul>