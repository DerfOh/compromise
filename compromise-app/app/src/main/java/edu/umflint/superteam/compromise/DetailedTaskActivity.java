package edu.umflint.superteam.compromise;

import android.content.SharedPreferences;
import android.os.Bundle;
import android.preference.PreferenceManager;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import java.security.acl.Group;
import java.util.ArrayList;

import edu.umflint.superteam.compromise.API.GetUsers;
import edu.umflint.superteam.compromise.API.NewTask;

public class DetailedTaskActivity extends AppCompatActivity {



    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_detailed_task);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

    final EditText points = (EditText) findViewById(R.id.Points);
    final EditText title = (EditText) findViewById(R.id.TaskTitle);
    final EditText desc = (EditText) findViewById(R.id.taskDescription);
    final Button createBtn = (Button) findViewById(R.id.createButton);

        Bundle bundle = getIntent().getExtras();
        points.setText(bundle.getString("taskPoints", ""));
        title.setText(bundle.getString("taskTitle", ""));
        desc.setText(bundle.getString("taskDescription", ""));

        if(bundle.getBoolean("newTask"))
        {
            createBtn.setVisibility(View.VISIBLE);

            createBtn.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
                    int GroupId = prefs.getInt("SelectedGroup", -1);
                    String response = "test";
                    try {
                        response = new NewTask(GroupId, title.getText().toString(), desc.getText().toString(), Integer.parseInt(points.getText().toString())).execute().get();
                    }catch (Exception e)
                    {
                        Toast.makeText(getApplicationContext(), "Task Created! " + e.toString(), Toast.LENGTH_SHORT).show();
                    }
                        Toast.makeText(getApplicationContext(), "Task Created! " + response, Toast.LENGTH_SHORT).show();
                }
            });


        } else{
            createBtn.setVisibility(View.INVISIBLE);
        }



        getSupportActionBar().setDisplayHomeAsUpEnabled(true);
    }

}
