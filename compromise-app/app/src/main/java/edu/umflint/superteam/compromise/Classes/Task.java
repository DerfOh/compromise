package edu.umflint.superteam.compromise.Classes;

import android.graphics.Point;

import java.security.acl.Group;

import edu.umflint.superteam.compromise.API.HttpClient;

/**
 * Created by nsammut on 3/26/17.
 */

public class Task {

    final private String url = "https://api.compromise.rocks/api/tasks/";

    private int TaskId;
    private int GroupId;
    private String Name;
    private String Description;
    private String CompletionStatus;
    private int PointValue;
    private String CompletedBy;

    public String getName() {
        return Name;
    }
    public void setName(String Name) {
        this.Name = Name;
    }
    public int getTaskId() {
        return TaskId;
    }
    public void setTaskId(int TaskId) {
        this.TaskId = TaskId;
    }
    public int getGroupId() {
        return GroupId;
    }
    public void setGroupId(int GroupId) {
        this.GroupId = GroupId;
    }
    public String getDescription() {
        return Description;
    }
    public void setDescription(String Description) {
        this.Description = Description;
    }
    public String getCompletionStatus() {
        return CompletionStatus;
    }
    public void setCompletionStatus(String CompletionStatus) {
        this.CompletionStatus = CompletionStatus;
    }
    public int getPointValue() {
        return PointValue;
    }
    public void setPointValue(int PointValue) {
        this.PointValue = PointValue;
    }
    public String getCompletedBy() {
        return CompletedBy;
    }
    public void setCompletedBy(String CompletedBy) {
        this.CompletedBy = CompletedBy;
    }

    public void create()
    {
        new HttpClient("POST", url, "GroupId=" + GroupId + "&TaskName=" + Name + "&TaskDescription=" + Description + "&CompletionStatus=Incomplete&PointValue=" + PointValue).execute();
    }
    public void modify()
    {
        new HttpClient("PUT", url, "GroupId=" + GroupId + "&TaskName=" + Name + "&TaskDescription=" + Description + "&CompletionStatus=" + CompletionStatus + "&PointValue=" + PointValue + "&CompletedBy=" + CompletedBy).execute();
    }
    public void complete()
    {
        new HttpClient("PUT", url, "TaskId=" + TaskId + "&GroupId=" + GroupId + "&TaskName=" + Name + "&TaskDescription=" + Description + "&CompletionStatus=Complete&CompletedBy=" + CompletedBy + "&PointValue=" + PointValue).execute();
    }
}
