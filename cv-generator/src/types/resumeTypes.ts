export interface Social {
    name: string;
    url: string;
}

export interface Link {
    label: string;
    href: string;
}

export interface Education {
    school: string;
    degree: string;
    start: string;
    end: string;
}

export interface WorkExperience {
    company: string;
    link: string;
    badges: string[];
    title: string;
    start: string;
    end: string;
    description: string;
}

export interface Project {
    title: string;
    link: Link;
    techStack: string[];
    description: string;
}

export interface Contact {
    email: string;
    tel: string;
    social: Social[];
}

export interface ResumeData {
    name: string;
    initials: string;
    location: string;
    locationLink: string;
    avatarUrl?: string;
    personalWebsiteUrl?: string;
    summary: string;
    about: string;
    contact: Contact;
    education: Education[];
    work: WorkExperience[];
    skills: string[];
    projects: Project[];
    hobbies: string[];
}