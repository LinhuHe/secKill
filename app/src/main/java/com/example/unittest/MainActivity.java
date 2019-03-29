package com.example.unittest;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

public class MainActivity extends AppCompatActivity {

    private EditText ed_textone;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ed_textone = (EditText)findViewById(R.id.ed_Textone);
    }

    public void OnJump(View view)
    {
        Toast.makeText(getApplication(),ed_textone.getText().toString(),Toast.LENGTH_SHORT).show();

        Intent intent = new Intent(getApplicationContext(),Activity_two.class);
        startActivity(intent);
    }
}
